module Main where


import Data.Char

readLines :: FilePath -> IO [String]
readLines = fmap lines . readFile

splitRows :: Int -> [a] -> [[a]]
splitRows n [] = []
splitRows n xs = take n xs:splitRows n (drop n xs)

getNumberOnSpace :: [[a]] -> Int -> [a]
getNumberOnSpace m k = map (!! k) m

splitBy :: Char -> String -> [String]
splitBy _ [] = []
splitBy c s  =
  let
    i = (length . takeWhile (/= c)) s
    (as, bs) = splitAt i s
  in as : splitBy c (if null bs then [] else tail bs)

createCorrectList :: [[String]] -> [[String]]
createCorrectList = map splitByComma

splitByComma :: [String] -> [String]
splitByComma [] = []
splitByComma(x:xs) = splitBy ',' x

toIntList :: [[String]] -> [[Int]]
toIntList = map toIntArray

toIntArray :: [String] -> [Int]
toIntArray = map (toInts . head)

toInts :: Char -> Int
toInts = digitToInt




main = do
  content <- readLines "PuzzleBoard.txt"
  let boardRows = splitRows 1 content
  let boardString = createCorrectList boardRows
  let board = toIntList boardString
  let currentSpace = [0,0]
  let endSpace = [9,9]
  let spacesVisited = []
  let badSpaces = []
  print board







--  let test =  board !! 2
--  let h = head test !! 1