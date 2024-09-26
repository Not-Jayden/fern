/**
 * This file was auto-generated by Fern from our API Definition.
 */
package com.seed.literal.model.reference;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.seed.literal.core.ObjectMappers;
import java.util.Objects;
import java.util.Optional;

@JsonInclude(JsonInclude.Include.NON_ABSENT)
@JsonDeserialize(builder = SendRequest.Builder.class)
public final class SendRequest {
    private final String query;

    private final String context;

    private final Optional<String> maybeContext;

    private final ContainerObject containerObject;

    private SendRequest(String query, String context, Optional<String> maybeContext, ContainerObject containerObject) {
        this.query = query;
        this.context = context;
        this.maybeContext = maybeContext;
        this.containerObject = containerObject;
    }

    @JsonProperty("prompt")
    public String getPrompt() {
        return "You are a helpful assistant";
    }

    @JsonProperty("query")
    public String getQuery() {
        return query;
    }

    @JsonProperty("stream")
    public Boolean getStream() {
        return false;
    }

    @JsonProperty("context")
    public String getContext() {
        return context;
    }

    @JsonProperty("maybeContext")
    public Optional<String> getMaybeContext() {
        return maybeContext;
    }

    @JsonProperty("containerObject")
    public ContainerObject getContainerObject() {
        return containerObject;
    }

    @java.lang.Override
    public boolean equals(Object other) {
        if (this == other) return true;
        return other instanceof SendRequest && equalTo((SendRequest) other);
    }

    private boolean equalTo(SendRequest other) {
        return query.equals(other.query)
                && context.equals(other.context)
                && maybeContext.equals(other.maybeContext)
                && containerObject.equals(other.containerObject);
    }

    @java.lang.Override
    public int hashCode() {
        return Objects.hash(this.query, this.context, this.maybeContext, this.containerObject);
    }

    @java.lang.Override
    public String toString() {
        return ObjectMappers.stringify(this);
    }

    public static QueryStage builder() {
        return new Builder();
    }

    public interface QueryStage {
        ContextStage query(String query);

        Builder from(SendRequest other);
    }

    public interface ContextStage {
        ContainerObjectStage context(String context);
    }

    public interface ContainerObjectStage {
        _FinalStage containerObject(ContainerObject containerObject);
    }

    public interface _FinalStage {
        SendRequest build();

        _FinalStage maybeContext(Optional<String> maybeContext);

        _FinalStage maybeContext(String maybeContext);
    }

    @JsonIgnoreProperties(ignoreUnknown = true)
    public static final class Builder implements QueryStage, ContextStage, ContainerObjectStage, _FinalStage {
        private String query;

        private String context;

        private ContainerObject containerObject;

        private Optional<String> maybeContext = Optional.empty();

        private Builder() {}

        @java.lang.Override
        public Builder from(SendRequest other) {
            query(other.getQuery());
            context(other.getContext());
            maybeContext(other.getMaybeContext());
            containerObject(other.getContainerObject());
            return this;
        }

        @java.lang.Override
        @JsonSetter("query")
        public ContextStage query(String query) {
            this.query = Objects.requireNonNull(query, "query must not be null");
            return this;
        }

        @java.lang.Override
        @JsonSetter("context")
        public ContainerObjectStage context(String context) {
            this.context = Objects.requireNonNull(context, "context must not be null");
            return this;
        }

        @java.lang.Override
        @JsonSetter("containerObject")
        public _FinalStage containerObject(ContainerObject containerObject) {
            this.containerObject = Objects.requireNonNull(containerObject, "containerObject must not be null");
            return this;
        }

        @java.lang.Override
        public _FinalStage maybeContext(String maybeContext) {
            this.maybeContext = Optional.ofNullable(maybeContext);
            return this;
        }

        @java.lang.Override
        @JsonSetter(value = "maybeContext", nulls = Nulls.SKIP)
        public _FinalStage maybeContext(Optional<String> maybeContext) {
            this.maybeContext = maybeContext;
            return this;
        }

        @java.lang.Override
        public SendRequest build() {
            return new SendRequest(query, context, maybeContext, containerObject);
        }
    }
}
