getwd()
setwd('C:/Users/GN4348/CODE/advent/2022/R/06')
lines <- readLines("06.txt")
lines[1]

codes <- as.list(strsplit(lines[1], "")[[1]]) 

data <-c()
for (i in 1:length(codes)) {
    data <- rbind(data,c(codes[i],codes[i+1],codes[i+2],codes[i+3]))
}
taskA <- match(4, apply(data,1,function(x) length(unique(x))))+3
taskA

data2 <-c()
for (i in 1:length(codes)) {
    data2 <- rbind(data2,c( codes[i],codes[i+1],codes[i+2],codes[i+3],
                            codes[i+4],codes[i+5],codes[i+6],codes[i+7],
                            codes[i+8],codes[i+9],codes[i+10],codes[i+11],
                            codes[i+12],codes[i+13]))
}
taskB <- match(14, apply(data2,1,function(x) length(unique(x))))+13
taskB