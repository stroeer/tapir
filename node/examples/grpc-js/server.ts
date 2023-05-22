import { createServer } from 'http';
import { getSectionPage } from './client';

const port = 3000;

const ignorePath = ['/favicon.ico', '/service-worker.js'];

const toResult = (data: unknown) => {
  return JSON.stringify(data, null, 2);
};

const server = createServer((req, res) => {
  if (ignorePath.includes(req.url ?? '')) {
    res.end();
  } else {
    getSectionPage(req.url ?? '/', (err, sectionPage) => {
      res.setHeader('Content-Type', 'application/javascript; charset=utf-8');
      if (err) {
        res.end(
          toResult({
            status: 'error',
            error: err,
          })
        );
      } else {
        res.end(
          toResult({
            status: 'success',
            path: sectionPage?.getSection()?.getFieldsMap().get('ref_path'),
          })
        );
      }
    });
  }
});

server.addListener('error', (err) => {
  console.error(err);
});

server.listen(port, () => {
  console.log(`server is listening on http://localhost:${port}`);
});
