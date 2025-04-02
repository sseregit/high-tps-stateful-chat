package schema

type Scannable interface {
	ScanRow(scanner interface{ Scan(dest ...any) error }) error
}
