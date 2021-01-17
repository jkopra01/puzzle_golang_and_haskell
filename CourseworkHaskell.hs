module Main where

readLines :: FilePath -> IO [String]
readLines = fmap lines . readFile

splitRows :: Int -> [a] -> [[a]]
splitRows n [] = []
splitRows n xs = (take n xs):(splitRows n (drop n xs))


main = do
  content <- readLines "PuzzleBoard.txt"
  let board = splitRows 1 content
  print board