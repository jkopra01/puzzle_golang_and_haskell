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

getOptionToAvoidNegIndU :: [Int] -> Int -> [Int]
getOptionToAvoidNegIndU x y = do
  if head x - y > 0
    then [head x - y, last x]
    else [0,0]

getOptionToAvoidNegIndL :: [Int] -> Int -> [Int]
getOptionToAvoidNegIndL x y = do
  if last x - y > 0
    then [head x, last x - y]
    else [0,0]



solve :: [[Int]] -> [Int] -> [[Int]] -> [[Int]] -> IO ()
solve board current visited bad = do
  print current
  let correctRow = getRow board current
  let numberOnSpace = getNumberOnSpace2 correctRow current
  let finalSpace = [9,9]
  let rightOption = [head current,last current + numberOnSpace]
  let downOption = [head current + numberOnSpace,last current]

  let leftOption = getOptionToAvoidNegIndL current numberOnSpace
  let upOption =  getOptionToAvoidNegIndU current numberOnSpace

  if current == finalSpace
    then print "ff"

  else if numberOnSpace+last current <= 9 && notElem rightOption bad && notElem rightOption visited 
                      then do 
                        let temp = visited ++ [rightOption]
                        solve board rightOption temp bad
  else if numberOnSpace+head current <= 9 && notElem downOption bad && notElem downOption visited 
                      then do 
                        let temp = visited ++ [downOption]
                        solve board downOption temp bad
  else if numberOnSpace-head current > 0 && notElem upOption bad && notElem upOption visited
                      then do 
                        let temp = visited ++ [upOption]
                        solve board upOption temp bad
  else if numberOnSpace-last current > 0 && notElem leftOption bad && notElem leftOption visited
                      then do 
                        let temp = visited ++ [leftOption]
                        solve board leftOption temp bad
  else do
    let bad = bad ++ [current]
    let visited = [[]]
    let current = [0,0]
    print bad
    solve board current visited bad

getRow :: [[Int]] -> [Int] -> [Int]
getRow b c = b !! head c

getNumberOnSpace2 :: [Int] ->  [Int] -> Int
getNumberOnSpace2 b c = b !! last c


main = do
  content <- readLines "PuzzleBoard.txt"
  let boardRows = splitRows 1 content
  let boardString = createCorrectList boardRows
  let board = toIntList boardString
  let currentSpace = [0,0]
 -- let endSpace = [9,9]
  let spacesVisited = [[]]
  let badSpaces = [[]]
 -- let correctSequence = solve board currentSpace spacesVisited badSpaces
 -- let correctSequence = getRow board currentSpace
 -- let correctSequence2 = getNumberOnSpace2 correctSequence currentSpace
  solve board currentSpace spacesVisited badSpaces
