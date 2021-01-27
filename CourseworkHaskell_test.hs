module Test where
import Test.HUnit
import System.Exit

import CourseworkHaskell as CH

boardFromFile = [["7","4","4","6","6","3","2","2","6","8"],["3","3","6","5","4","3","7","2","8","3"],["4","1","6","6","2","4","4","4","7","4"],["4","5","3","4","3","5","4","4","8","5"],["5","1","4","6","6","5","0","7","1","4"],["2","6","9","4","9","7","7","9","1","4"],["3","5","4","0","6","4","5","5","5","6"],["6","6","2","3","4","7","1","2","3","3"],["3","5","4","3","6","5","4","5","2","6"],["3","9","3","5","1","1","5","4","6","0"]]
boardFromPic = [[7,4,4,6,6,3,2,2,6,8],[3,3,6,5,4,3,7,2,8,3],[4,1,6,6,2,4,4,4,7,4],[4,5,3,4,3,5,4,4,8,5],[5,1,4,6,6,5,0,7,1,4],[2,6,9,4,9,7,7,9,1,4],[3,5,4,0,6,4,5,5,5,6],[6,6,2,3,4,7,1,2,3,3],[3,5,4,3,6,5,4,5,2,6],[3,9,3,5,1,1,5,4,6,0]]
correctPath = [[],[0,7],[2,7],[6,7],[1,7],[1,5],[1,2],[7,2],[7,4],[7,8],[4,8],[5,8],[5,9],[9,9]]
row = [4,5,3,4,3,5,4,4,8,5]
space =  [3,2]

firstSpace = [0,0]
spacesVisited = [[]]
badSpaces = [[]]

testBoardCase = TestCase(assertEqual "Board is correct" boardFromPic (CH.toIntList boardFromFile))
testSolveCase = TestCase(assertEqual "Puzzle gets solved" correctPath (CH.solve boardFromPic firstSpace spacesVisited badSpaces))
testRowCase = TestCase(assertEqual "Correct row" (row) (CH.getRow boardFromPic space))
testNumberInSpaceCase = TestCase(assertEqual "Correct number from space" 3 (CH.getNumberOnSpace row space))

main :: IO ()
main = do
    counts <- runTestTT ( test [
        testBoardCase,
        testSolveCase,
        testRowCase,
        testNumberInSpaceCase
        ])
    if errors counts + failures counts == 0
        then exitSuccess
        else exitFailure