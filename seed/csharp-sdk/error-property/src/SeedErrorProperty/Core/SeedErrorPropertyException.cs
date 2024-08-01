using System;

#nullable enable

namespace SeedErrorProperty.Core;

/// <summary>
/// Base exception class for all exceptions thrown by the SDK.
/// </summary>
public class SeedErrorPropertyException(string message, Exception? innerException = null)
    : Exception(message, innerException) { }
