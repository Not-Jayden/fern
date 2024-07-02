using System.Text.Json.Serialization;

#nullable enable

namespace SeedTrace.V2.V3;

public class LightweightProblemInfoV2
{
    [JsonPropertyName("problemId")]
    public string ProblemId { get; init; }

    [JsonPropertyName("problemName")]
    public string ProblemName { get; init; }

    [JsonPropertyName("problemVersion")]
    public int ProblemVersion { get; init; }

    [JsonPropertyName("variableTypes")]
    public HashSet<object> VariableTypes { get; init; }
}
