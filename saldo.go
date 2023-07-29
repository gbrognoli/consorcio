package main

import "fmt"

func main() {
    // Informações do consórcio
    var numeroTotalPrestacoes, prestacoesPagas, valorPrestacao int

    // Leitura dos dados do consórcio
    fmt.Print("Informe o número total de prestações: ")
    fmt.Scanln(&numeroTotalPrestacoes)

    fmt.Print("Informe a quantidade total de prestações pagas: ")
    fmt.Scanln(&prestacoesPagas)

    fmt.Print("Informe o valor de cada prestação: ")
    fmt.Scanln(&valorPrestacao)

    // Cálculo do total pago pelo consorciado
    totalPago := prestacoesPagas * valorPrestacao

    // Cálculo do saldo devedor
    saldoDevedor := (numeroTotalPrestacoes - prestacoesPagas) * valorPrestacao

    // Exibindo os resultados
    fmt.Printf("Total pago pelo consorciado: %d\n", totalPago)
    fmt.Printf("Saldo devedor: %d\n", saldoDevedor)
}
