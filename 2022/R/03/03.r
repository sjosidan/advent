setwd("C:/Users/sjosi/CODE/advent/2022/R/03")
lines <- readLines("03.txt")
alphabet <- c(letters, LETTERS)

halves <- lapply(lines, function(x) {
  n <- nchar(x)
  half1 <- substr(x, 1, n/2)
  half2 <- substr(x, n/2+1, n)
  return(c(half1, half2))
})

task03a <- 0 
for (item in halves) {
    string1 <- item[1]
    string2 <- item[2]
    characters1 <- unlist(strsplit(string1, ""))
    characters2 <- unlist(strsplit(string2, ""))
    intersections <- intersect(characters1, characters2)

    task03a = task03a + match(intersections[1], alphabet)
}

task03b <- 0 
localCounter <- 1
lastChars <- list()
for (l in lines) {
    characters <- unlist(strsplit(l, ""))
    if (localCounter == 1) {
        lastChars = characters
    }
    intersections <- intersect(characters, lastChars)
    if (localCounter == 3) {
        task03b = task03b + match(intersections[1], alphabet)
        localCounter = 0
    }
    lastChars <- intersections
    localCounter = localCounter + 1
}

task03a
task03b