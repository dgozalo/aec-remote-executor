public class TestsRunner {
    public static void main(String[] args) {
        if (args.length == 0) {
            throw new RuntimeException("Test runner requires a path to a file with test cases");
        }

        String testsFilePath = args[0];
        if (testsFilePath == null || testsFilePath.isEmpty()) {
            throw new RuntimeException("The path to the test file is empty or null");
        }

        TestsRunner runner = new TestsRunner();
        Solution solution = new Solution();

        for (int i = 1; i <= 10; i++) {
            int[] sol = solution.twoSum(new int[0], i);
            runner.writeResultToTestsFile(testsFilePath, String.format("TestCase#%1d::%2s::%3s::%4b", i,
                    java.util.Arrays.toString(sol),
                    java.util.Arrays.toString(new int[]{1, 2}),
                    java.util.Arrays.equals(sol, new int[]{1, 2})));
        }
    }

    public void writeResultToTestsFile(String filePath, String result) {
        //write to a file
        java.io.File testsFile = new java.io.File(filePath);
        if (!testsFile.exists()) {
            throw new RuntimeException("Test File does not exist");
        }
        if (!testsFile.canWrite()) {
            throw new RuntimeException("Test File is not writable");
        }
        try (java.io.FileWriter fileWriter = new java.io.FileWriter(testsFile, true);
             java.io.BufferedWriter bufferedWriter = new java.io.BufferedWriter(fileWriter)) {
            bufferedWriter.write(result);
            bufferedWriter.newLine();
        } catch (java.io.IOException e) {
            throw new RuntimeException("Error writing to file", e);
        }
    }
}

class Solution {
    public int[] twoSum(int[] nums, int target) {
        System.out.println("test");
        return new int[]{target};
    }
}

