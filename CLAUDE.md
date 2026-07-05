# CLAUDE.md — theia-subtui

Fork de [MattiaPun/SubTUI](https://github.com/MattiaPun/SubTUI) mantenido por Rodrigo Mera.
Cliente TUI (terminal) para Navidrome/Subsonic. Lenguaje: Go. UI: Bubble Tea + Lipgloss.

## Regla obligatoria — CHANGELOG

**Todo cambio al código DEBE registrarse en `CHANGELOG.md` antes de declarar la tarea terminada.**

- Sección activa: `## [Unreleased]`
- Subsecciones: `### Fixed`, `### Changed`, `### Added`, `### Removed`
- Formato de entrada: descripción concisa + archivo(s) afectado(s) + motivo
- Al crear una release (`theia-X.Y.Z`): mover entradas de `[Unreleased]` a la nueva sección con fecha

No hay excepción. Un cambio sin entrada en CHANGELOG = tarea incompleta.

## Build y deploy

```bash
# Compilar
cd ~/projects/theia-subtui
go build -o theia-subtui .

# Instalar (reemplaza el binario del sistema)
sudo cp theia-subtui /usr/bin/subtui

# Verificar versión (debe mostrar "| theia-subtui" al final)
subtui -v
```

## Convención de versiones

`theia-X.Y.Z` — ver `CHANGELOG.md` para detalles.
Upstream base actual: v2.14.3 (commit e5702ad).

## Estructura relevante

```
internal/ui/view.go   — toda la lógica de renderizado (Bubble Tea View)
internal/ui/update.go — manejo de eventos (Bubble Tea Update)
main.go               — entry point, flags, versión
CHANGELOG.md          — historial de cambios del fork (OBLIGATORIO actualizar)
```

## Pitfalls conocidos

- `strings.Repeat` con count negativo causa panic — siempre guardar con `if n < 0 { n = 0 }` antes de llamar
- El guard `if m.width < 50 || m.height < 25` en `View()` no protege renders intermedios durante SIGWINCH
- `calculateMainWidthAndHeight` devuelve `availableWidth` que puede ser negativo si el terminal es muy angosto

## Git

- Commits en español o inglés, cortos y descriptivos
- No incluir referencias a IA en commits ni PRs
- Push directo a `main` (repo personal, sin PR salvo colaboración externa)
