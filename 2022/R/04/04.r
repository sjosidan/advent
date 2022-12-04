setwd("C:/Users/sjosi/CODE/advent/2022/R/04")
data <- read.csv("04.csv", sep=";")
n <- nrow(data)
task04a <- 0
task04b <- 0
for (i in 1:n) {
    range1 <- seq(from = data[i, 1], to = data[i, 2], by = 1)
    range2 <- seq(from = data[i, 3], to = data[i, 4], by = 1)

    if (all(range2 %in% range1)) {
        task04a <- task04a + 1
    } else if (all(range1 %in% range2)) {
        task04a <- task04a + 1
    }

    if (any(range2 %in% range1)) {
        task04b <- task04b + 1
    } else if (any(range1 %in% range2)) {
        task04b <- task04b + 1
    }
}

task04a
task04b