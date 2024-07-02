# trace-logger

Логер с поддержкой opentelemetrecy

### Примеры
Примеры использвония логера

`rootLogger, err := logger.New(cfg)` – Создание корнеого логера

`log, parentCtx := m.log.Named(ctx, "SpanName")` – Создание логера наследуемого от ctx логера. Для создания дочернего логера надо передать контекст родительского

Пример:

`log, parentCtx := m.log.Named(ctx, "ParenSpan")` 
`log, childCtx := m.log.Named(ctx, "ChildSpan")` 

Интерфейс логирования аалогичный zap но с передачей контеста.
`log.Debug(ctx, "get response", zap.Any("response", resp))`
