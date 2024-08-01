using System;

#nullable enable

namespace SeedMultiUrlEnvironment.Core;

/// <summary>
/// Base exception class for all exceptions thrown by the SDK.
/// </summary>
public class SeedMultiUrlEnvironmentException(string message, Exception? innerException = null)
    : Exception(message, innerException) { }
