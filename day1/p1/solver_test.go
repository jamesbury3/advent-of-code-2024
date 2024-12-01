package d1p1

import "testing"

func TestD1P1_Solve(t *testing.T) {
	type fields struct {
		lines []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Solve Problem",
			want: "936063",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := &D1P1{
				lines: tt.fields.lines,
			}
			got, err := solver.Solve()
			if (err != nil) != tt.wantErr {
				t.Errorf("D1P1.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("D1P1.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
