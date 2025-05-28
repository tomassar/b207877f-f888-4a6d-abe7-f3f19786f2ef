# Portfolio Rebalancer

This is a simple Go application that helps you rebalance your investment portfolio. You define your current stock holdings and your target allocation percentages, and it tells you exactly what to buy or sell to match your goals. For example, if you want 60% Apple and 40% Meta, but you're currently overweight in Meta, it'll suggest selling some Meta shares and buying more Apple.

To run it, just clone the repo and run `go run main.go`. The example shows a portfolio with Apple and Meta stocks that gets rebalanced according to a 60/40 allocation. You can easily modify the stocks, shares, prices, and target allocations in the main.go file to match your own portfolio. 