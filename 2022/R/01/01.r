setwd("C:/Users/sjosi/CODE/advent/2022/R/01")
my_data <- read.csv("01.csv", sep=';')
sorted = rowSums(my_data, na.rm=TRUE)
taskA = max(sorted)
n <- nrow(my_data)
taskB <- sort(sorted,partial=n-1)[n-1] + sort(sorted,partial=n-2)[n-2] + taskA
taskA
taskB

