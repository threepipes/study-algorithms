package parser

import (
	"calculator/ast"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name    string
		args    args
		want    *ast.Expression
		wantErr bool
	}{
		{
			name: "simple add",
			args: args{
				tokens: []string{"1", "+", "2"},
			},
			want: &ast.Expression{
				Ts: []*ast.Term{
					{
						Fs: []*ast.Factor{
							{
								Num: &ast.Number{Value: 1},
							},
						},
						Ops: []*ast.OperatorMultLike{nil},
					},
					{
						Fs: []*ast.Factor{
							{
								Num: &ast.Number{Value: 2},
							},
						},
						Ops: []*ast.OperatorMultLike{nil},
					},
				},
				Ops: []*ast.OperatorAddLike{
					{Op: ast.OperatorAddLikeKindAdd},
					{Op: ast.OperatorAddLikeKindAdd},
				},
			},
		},
		{
			name: "simple subtract",
			args: args{
				tokens: []string{"1", "-", "2"},
			},
			want: &ast.Expression{
				Ts: []*ast.Term{
					{
						Fs: []*ast.Factor{
							{
								Num: &ast.Number{Value: 1},
							},
						},
						Ops: []*ast.OperatorMultLike{nil},
					},
					{
						Fs: []*ast.Factor{
							{
								Num: &ast.Number{Value: 2},
							},
						},
						Ops: []*ast.OperatorMultLike{nil},
					},
				},
				Ops: []*ast.OperatorAddLike{
					{Op: ast.OperatorAddLikeKindAdd},
					{Op: ast.OperatorAddLikeKindSub},
				},
			},
		},
		{
			name: "mult and add",
			args: args{
				tokens: []string{"1", "+", "2", "*", "3"},
			},
			want: &ast.Expression{
				Ts: []*ast.Term{
					{
						Fs: []*ast.Factor{
							{
								Num: &ast.Number{Value: 1},
							},
						},
						Ops: []*ast.OperatorMultLike{nil},
					},
					{
						Fs: []*ast.Factor{
							{
								Num: &ast.Number{Value: 2},
							},
							{
								Num: &ast.Number{Value: 3},
							},
						},
						Ops: []*ast.OperatorMultLike{
							nil,
							{Op: ast.OperatorMultLikeKindMult},
						},
					},
				},
				Ops: []*ast.OperatorAddLike{
					{Op: ast.OperatorAddLikeKindAdd},
					{Op: ast.OperatorAddLikeKindAdd},
				},
			},
		},
		{
			name: "include bracket",
			args: args{
				tokens: []string{"1", "/", "(", "2", "-", "1", ")"},
			},
			want: &ast.Expression{
				Ts: []*ast.Term{
					{
						Fs: []*ast.Factor{
							{
								Num: &ast.Number{Value: 1},
							},
							{
								Exp: &ast.Expression{
									Ts: []*ast.Term{
										{
											Fs: []*ast.Factor{
												{
													Num: &ast.Number{Value: 2},
												},
											},
											Ops: []*ast.OperatorMultLike{nil},
										},
										{
											Fs: []*ast.Factor{
												{
													Num: &ast.Number{Value: 1},
												},
											},
											Ops: []*ast.OperatorMultLike{nil},
										},
									},
									Ops: []*ast.OperatorAddLike{
										{Op: ast.OperatorAddLikeKindAdd},
										{Op: ast.OperatorAddLikeKindSub},
									},
								},
							},
						},
						Ops: []*ast.OperatorMultLike{
							nil,
							{Op: ast.OperatorMultLikeKindDiv},
						},
					},
				},
				Ops: []*ast.OperatorAddLike{
					{Op: ast.OperatorAddLikeKindAdd},
				},
			},
		},
		{
			name: "include a minus operator at the top of the ast.Expression",
			args: args{
				tokens: []string{"-", "1", "+", "2"},
			},
			want: &ast.Expression{
				Ts: []*ast.Term{
					{
						Fs: []*ast.Factor{
							{
								Num: &ast.Number{Value: 1},
							},
						},
						Ops: []*ast.OperatorMultLike{nil},
					},
					{
						Fs: []*ast.Factor{
							{
								Num: &ast.Number{Value: 2},
							},
						},
						Ops: []*ast.OperatorMultLike{nil},
					},
				},
				Ops: []*ast.OperatorAddLike{
					{Op: ast.OperatorAddLikeKindSub},
					{Op: ast.OperatorAddLikeKindAdd},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.tokens)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
