using System;

#nullable enable

namespace SeedAliasExtends.Core;

/// <summary>
/// Base exception class for all exceptions thrown by the SDK.
/// </summary>
public class SeedAliasExtendsException(string message, Exception? innerException = null)
    : Exception(message, innerException) { }
