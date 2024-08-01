using System;

#nullable enable

namespace SeedPagination.Core;

/// <summary>
/// Base exception class for all exceptions thrown by the SDK.
/// </summary>
public class SeedPaginationException(string message, Exception? innerException = null)
    : Exception(message, innerException) { }
