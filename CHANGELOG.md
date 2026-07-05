# Changelog — theia-subtui

Fork de [MattiaPun/SubTUI](https://github.com/MattiaPun/SubTUI) mantenido por Rodrigo Mera.
Upstream base: v2.14.3 (commit e5702ad).

---

## [Unreleased]

## [theia-1.0.0] — 2026-07-05

### Forkeado desde upstream v2.14.3

### Fixed
- **Panic en cambio de escritorio virtual (SIGWINCH):** `strings.Repeat` en `footerInformation` recibía count negativo cuando el terminal reportaba un ancho reducido al cambiar de workspace en GNOME. Fix: guard `if progressLen < 0` y `if dashLen < 0` antes de cada `strings.Repeat` en la barra de progreso del footer (`internal/ui/view.go`).
- **`calculateMainWidthAndHeight` sin guard de ancho:** la función guardaba `availableHeight < 1` pero no `availableWidth`. Fix: `if availableWidth < 0 { availableWidth = 0 }`.

### Changed
- **Versión con distintivo:** `subtui -v` muestra `| theia-subtui` al final para identificar el fork vs. el binario upstream.
- **Tag visual en UI:** borde superior del panel principal muestra ` theia ` (visible en todo momento, con o sin cola activa). Implementado en `internal/ui/view.go` función `footerStyle()`.

---

## Convención de versiones

`theia-X.Y.Z` donde:
- `X` = versión mayor upstream en la que se basa
- `Y` = releases propios con cambios significativos
- `Z` = fixes/patches menores

Upstream merges: registrar en una entrada `[theia-X.Y.Z] — merged upstream vA.B.C`.
