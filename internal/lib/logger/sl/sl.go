package sl

import (
    "log/slog"

    // "url-shortener/internal/lib/logger/handlers/slogdiscard"
)

func Err(err error) slog.Attr {
    return slog.Attr{
        Key:   "error",
        Value: slog.StringValue(err.Error()),
    }
}