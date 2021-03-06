package pnc

import (
  "math/big"
)

type BlumMicali struct {
  P *big.Int
  G *big.Int
  Term *big.Int
}

func NewBlumMicali() BlumMicali {
  b := BlumMicali{}
  return b
}

// p: large prime
// g: primitive root modulo p
func (b *BlumMicali) Seed(p *big.Int, g *big.Int, s *big.Int) {
  b.P = p
  b.G = g
  b.Term = s
}

func (b *BlumMicali) generate_next_term() {
  b.Term.Exp(b.G, b.Term, b.P)
}

func (b *BlumMicali) Bit() uint32 {
  b.generate_next_term()

  numerator := big.NewInt(1)
  numerator.Sub(b.P, numerator)
  denominator := big.NewInt(2)

  n := big.NewInt(0)
  n.Div(numerator, denominator)

  if b.Term.Cmp(n) == -1 {
    return 1
  }
  return 0
}

func (b *BlumMicali) Urand32() uint32 {
  urand32 := uint32(0)
  for i := 0; i < 32; i++ {
    urand32 = urand32 << 1 + b.Bit()
  }
  return urand32
}
