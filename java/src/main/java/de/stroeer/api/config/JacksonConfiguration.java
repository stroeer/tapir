package de.stroeer.api.config;

import static java.util.Collections.singletonList;

import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.module.SimpleDeserializers;
import com.fasterxml.jackson.databind.module.SimpleModule;
import com.fasterxml.jackson.databind.Module;
import com.fasterxml.jackson.databind.module.SimpleSerializers;
import com.google.protobuf.GeneratedMessageV3;
import java.util.Map;
import java.util.stream.Collectors;
import org.reflections.Reflections;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

// borrowed from https://github.com/venth/training-webflux-grpc
@Configuration
public class JacksonConfiguration {

  @Bean
  public Module protobufModule() {
    SimpleModule module = new SimpleModule("protobufModule");

    Map<Class<?>, JsonDeserializer<?>> deserializers = new Reflections("de.stroeer")
        .getSubTypesOf(GeneratedMessageV3.class)
        .stream()
        .collect(Collectors.toMap(clazz -> clazz, ProtobufDeserializer::new));

    module.setDeserializers(new SimpleDeserializers(deserializers));
    module.setSerializers(new SimpleSerializers(singletonList(new ProtobufSerializer())));

    return module;
  }
}