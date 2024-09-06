/**
 * This file was auto-generated by Fern from our API Definition.
 */
package com.seed.packageYml.types;

import com.fasterxml.jackson.annotation.JsonAnyGetter;
import com.fasterxml.jackson.annotation.JsonAnySetter;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.seed.packageYml.core.ObjectMappers;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
import org.jetbrains.annotations.NotNull;

@JsonInclude(JsonInclude.Include.NON_ABSENT)
@JsonDeserialize(builder = EchoRequest.Builder.class)
public final class EchoRequest {
    private final String name;

    private final int size;

    private final Map<String, Object> additionalProperties;

    private EchoRequest(String name, int size, Map<String, Object> additionalProperties) {
        this.name = name;
        this.size = size;
        this.additionalProperties = additionalProperties;
    }

    @JsonProperty("name")
    public String getName() {
        return name;
    }

    @JsonProperty("size")
    public int getSize() {
        return size;
    }

    @java.lang.Override
    public boolean equals(Object other) {
        if (this == other) return true;
        return other instanceof EchoRequest && equalTo((EchoRequest) other);
    }

    @JsonAnyGetter
    public Map<String, Object> getAdditionalProperties() {
        return this.additionalProperties;
    }

    private boolean equalTo(EchoRequest other) {
        return name.equals(other.name) && size == other.size;
    }

    @java.lang.Override
    public int hashCode() {
        return Objects.hash(this.name, this.size);
    }

    @java.lang.Override
    public String toString() {
        return ObjectMappers.stringify(this);
    }

    public static NameStage builder() {
        return new Builder();
    }

    public interface NameStage {
        SizeStage name(@NotNull String name);

        Builder from(EchoRequest other);
    }

    public interface SizeStage {
        _FinalStage size(int size);
    }

    public interface _FinalStage {
        EchoRequest build();
    }

    @JsonIgnoreProperties(ignoreUnknown = true)
    public static final class Builder implements NameStage, SizeStage, _FinalStage {
        private String name;

        private int size;

        @JsonAnySetter
        private Map<String, Object> additionalProperties = new HashMap<>();

        private Builder() {}

        @java.lang.Override
        public Builder from(EchoRequest other) {
            name(other.getName());
            size(other.getSize());
            return this;
        }

        @java.lang.Override
        @JsonSetter("name")
        public SizeStage name(@NotNull String name) {
            this.name = Objects.requireNonNull(name, "name must not be null");
            return this;
        }

        @java.lang.Override
        @JsonSetter("size")
        public _FinalStage size(int size) {
            this.size = size;
            return this;
        }

        @java.lang.Override
        public EchoRequest build() {
            return new EchoRequest(name, size, additionalProperties);
        }
    }
}
