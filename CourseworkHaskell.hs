module Main where
import Data.Char
import HUnit

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

--to avoid the error given by the negative index
getOptionToAvoidNegIndU :: [Int] -> Int -> [Int]
getOptionToAvoidNegIndU x y = 
  if head x - y > -1
    then [head x - y, last x]
    else [400,400]

getOptionToAvoidNegIndL :: [Int] -> Int -> [Int]
getOptionToAvoidNegIndL x y =
  if last x - y > -1
    then [head x, last x - y]
    else [400,400]



solve :: [[Int]] -> [Int] -> [[Int]] -> [[Int]] -> [[Int]]
solve board current visited bad = do
  let correctRow = getRow board current
  let numberOnSpace = getNumberOnSpace2 correctRow current
  let finalSpace = [9,9]
  let rightOption = [head current,last current + numberOnSpace]
  let downOption = [head current + numberOnSpace,last current]

  let leftOption = getOptionToAvoidNegIndL current numberOnSpace
  let upOption =  getOptionToAvoidNegIndU current numberOnSpace

  if current == finalSpace
    then visited

  else if numberOnSpace+last current < 10  && notElem rightOption bad && notElem rightOption visited 
                      then do 
                        let tempVisited = visited ++ [rightOption]
                        solve board rightOption tempVisited bad
  else if numberOnSpace+head current < 10 && notElem downOption bad && notElem downOption visited 
                      then do 
                        let tempVisited = visited ++ [downOption]
                        solve board downOption tempVisited bad
  else if head current - numberOnSpace > -1 && notElem upOption bad && notElem upOption visited && head upOption < 11
                      then do 
                        let tempVisited = visited ++ [upOption]
                        solve board upOption tempVisited bad
  else if last current - numberOnSpace > -1 && notElem leftOption bad && notElem leftOption visited && head leftOption < 11
                      then do 
                        let tempVisited = visited ++ [leftOption]
                        solve board leftOption tempVisited bad
  else 
                      do
                        let tempBad = bad ++ [current]
                        let visited = [[]]
                        let current = [0,0]
                        solve board current visited tempBad

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
  let spacesVisited = [[]]
  let badSpaces = [[]]
  let correctSequence = solve board currentSpace spacesVisited badSpaces
  print correctSequence