/**
 * This file was auto-generated by Fern from our API Definition.
 */
package com.seed.examples.resources.types.types;

import com.fasterxml.jackson.annotation.JsonAnyGetter;
import com.fasterxml.jackson.annotation.JsonAnySetter;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.seed.examples.core.ObjectMappers;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
import org.jetbrains.annotations.NotNull;

@JsonInclude(JsonInclude.Include.NON_ABSENT)
@JsonDeserialize(builder = Request.Builder.class)
public final class Request {
    private final Object request;

    private final Map<String, Object> additionalProperties;

    private Request(Object request, Map<String, Object> additionalProperties) {
        this.request = request;
        this.additionalProperties = additionalProperties;
    }

    @JsonProperty("request")
    public Object getRequest() {
        return request;
    }

    @java.lang.Override
    public boolean equals(Object other) {
        if (this == other) return true;
        return other instanceof Request && equalTo((Request) other);
    }

    @JsonAnyGetter
    public Map<String, Object> getAdditionalProperties() {
        return this.additionalProperties;
    }

    private boolean equalTo(Request other) {
        return request.equals(other.request);
    }

    @java.lang.Override
    public int hashCode() {
        return Objects.hash(this.request);
    }

    @java.lang.Override
    public String toString() {
        return ObjectMappers.stringify(this);
    }

    public static RequestStage builder() {
        return new Builder();
    }

    public interface RequestStage {
        _FinalStage request(@NotNull Object request);

        Builder from(Request other);
    }

    public interface _FinalStage {
        Request build();
    }

    @JsonIgnoreProperties(ignoreUnknown = true)
    public static final class Builder implements RequestStage, _FinalStage {
        private Object request;

        @JsonAnySetter
        private Map<String, Object> additionalProperties = new HashMap<>();

        private Builder() {}

        @java.lang.Override
        public Builder from(Request other) {
            request(other.getRequest());
            return this;
        }

        @java.lang.Override
        @JsonSetter("request")
        public _FinalStage request(@NotNull Object request) {
            this.request = Objects.requireNonNull(request, "request must not be null");
            return this;
        }

        @java.lang.Override
        public Request build() {
            return new Request(request, additionalProperties);
        }
    }
}
