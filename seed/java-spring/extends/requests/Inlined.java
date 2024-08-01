/**
 * This file was auto-generated by Fern from our API Definition.
 */

package requests;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import core.ObjectMappers;
import java.lang.Object;
import java.lang.String;
import java.util.Objects;
import types.IExampleType;

@JsonInclude(JsonInclude.Include.NON_ABSENT)
@JsonDeserialize(
    builder = Inlined.Builder.class
)
public final class Inlined implements IExampleType {
  private final String name;

  private final String docs;

  private final String unique;

  private Inlined(String name, String docs, String unique) {
    this.name = name;
    this.docs = docs;
    this.unique = unique;
  }

  @JsonProperty("name")
  @java.lang.Override
  public String getName() {
    return name;
  }

  @JsonProperty("docs")
  @java.lang.Override
  public String getDocs() {
    return docs;
  }

  @JsonProperty("unique")
  public String getUnique() {
    return unique;
  }

  @java.lang.Override
  public boolean equals(Object other) {
    if (this == other) return true;
    return other instanceof Inlined && equalTo((Inlined) other);
  }

  private boolean equalTo(Inlined other) {
    return name.equals(other.name) && docs.equals(other.docs) && unique.equals(other.unique);
  }

  @java.lang.Override
  public int hashCode() {
    return Objects.hash(this.name, this.docs, this.unique);
  }

  @java.lang.Override
  public String toString() {
    return ObjectMappers.stringify(this);
  }

  public static NameStage builder() {
    return new Builder();
  }

  public interface NameStage {
    DocsStage name(String name);

    Builder from(Inlined other);
  }

  public interface DocsStage {
    UniqueStage docs(String docs);
  }

  public interface UniqueStage {
    _FinalStage unique(String unique);
  }

  public interface _FinalStage {
    Inlined build();
  }

  @JsonIgnoreProperties(
      ignoreUnknown = true
  )
  public static final class Builder implements NameStage, DocsStage, UniqueStage, _FinalStage {
    private String name;

    private String docs;

    private String unique;

    private Builder() {
    }

    @java.lang.Override
    public Builder from(Inlined other) {
      name(other.getName());
      docs(other.getDocs());
      unique(other.getUnique());
      return this;
    }

    @java.lang.Override
    @JsonSetter("name")
    public DocsStage name(String name) {
      this.name = name;
      return this;
    }

    @java.lang.Override
    @JsonSetter("docs")
    public UniqueStage docs(String docs) {
      this.docs = docs;
      return this;
    }

    @java.lang.Override
    @JsonSetter("unique")
    public _FinalStage unique(String unique) {
      this.unique = unique;
      return this;
    }

    @java.lang.Override
    public Inlined build() {
      return new Inlined(name, docs, unique);
    }
  }
}
