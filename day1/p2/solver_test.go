package d1p2

import "testing"

func TestD1P2_Solve(t *testing.T) {
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
			want: "23150395",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := &D1P2{
				lines: tt.fields.lines,
			}
			got, err := solver.Solve()
			if (err != nil) != tt.wantErr {
				t.Errorf("D1P2.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("D1P2.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
