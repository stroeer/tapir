package de.stroeer.api.config;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.SerializerProvider;
import com.fasterxml.jackson.databind.ser.std.StdSerializer;
import com.google.protobuf.GeneratedMessageV3;
import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.util.JsonFormat;
import java.io.IOException;

// borrowed from https://github.com/venth/training-webflux-grpc
public class ProtobufSerializer extends StdSerializer<GeneratedMessageV3> {

  private final JsonFormat.Printer printer;

  public ProtobufSerializer() {
    super(GeneratedMessageV3.class);
    printer = JsonFormat.printer()
        .preservingProtoFieldNames()
        .omittingInsignificantWhitespace();
  }

  @Override
  public Class<GeneratedMessageV3> handledType() {
    return GeneratedMessageV3.class;
  }

  @Override
  public void serialize(GeneratedMessageV3 value, JsonGenerator gen, SerializerProvider serializers)
      throws IOException, InvalidProtocolBufferException {
    gen.writeRawValue(printer.print(value));
  }
}