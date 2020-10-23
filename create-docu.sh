export RELEASE_TAG=0.1.2
cd tapir.wiki
[ -d api-specs/$RELEASE_TAG ] || mkdir -p api-specs/$RELEASE_TAG
yes | cp ../api-docs/api-spec* ./api-specs/$RELEASE_TAG
cd ./api-specs/$RELEASE_TAG
echo "\n\n## $RELEASE_TAG \n" >> ../../api-spec.md
for n in api-spec-stroeer*; do filepath=${n/.md/}; label=${filepath/api-spec-/}; echo "• [$label](api-specs/$RELEASE_TAG/$filepath)" >> ../../api-spec.md; done
cd ../..
git diff api-spec.md
cat api-spec.md
rm -rf api-specs/*
git restore api-specs
git restore api-spec.md
cd ..
