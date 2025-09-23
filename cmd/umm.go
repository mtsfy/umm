package cmd

func init() {
	rootCmd.AddCommand(plusCmd)
	rootCmd.AddCommand(historyCmd)

	historyCmd.Flags().Int("page", -1, "page number for pagination")
	historyCmd.Flags().Int("size", -1, "number of interactions to display per page")
	historyCmd.Flags().String("search", "", "filter history by keywords in queries and responses")
	historyCmd.Flags().Int("delete", -1, "delete the specified history chat by its id")
}
