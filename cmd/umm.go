package cmd

func init() {
	rootCmd.AddCommand(plusCmd)
	rootCmd.AddCommand(historyCmd)

	historyCmd.Flags().IntP("page", "p", -1, "page number for pagination")
	historyCmd.Flags().IntP("size", "s", -1, "number of interactions to display per page")
	historyCmd.Flags().String("search", "", "filter history by keywords in queries and responses")
}
