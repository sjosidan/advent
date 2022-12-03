setwd("C:/Users/sjosi/CODE/advent/2022/R/02")
my_data <- read.csv("02.csv", sep=';')
head(my_data)
task02a <- sum(rowSums(my_data[ , c(3,4)], na.rm=TRUE))
task02b <- sum(rowSums(my_data[ , c(5,6)], na.rm=TRUE))
task02a
task02b