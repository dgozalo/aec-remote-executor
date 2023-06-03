/* eslint-disable react/prop-types */
/* eslint-disable react/function-component-definition */
/**
 =========================================================
 * Material Dashboard 2 React - v2.2.0
 =========================================================

 * Product Page: https://www.creative-tim.com/product/material-dashboard-react
 * Copyright 2023 Creative Tim (https://www.creative-tim.com)

 Coded by www.creative-tim.com

 =========================================================

 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 */

// @mui material components
import MDBox from "components/MDBox";
import TimelineItem from "examples/Timeline/TimelineItem";


export default function buildTestResultsTable(testResults) {
    return {
        columns: [
            {Header: "Test Name", accessor: "testName", align: "left"},
            {Header: "Expected Output", accessor: "expected", align: "center"},
            {Header: "Actual Output", accessor: "actual", align: "center"},
            {Header: "Pass", accessor: "passed", align: "center"}
        ],
        rows: testResults.map((testResult) => (
            {
                testName: testResult.testName,
                expected: testResult.expected,
                actual: testResult.actual,
                passed: (
                    <MDBox width="2rem" textAlign="right">
                        <TimelineItem
                            color={testResult.passed ? "success" : "error"}
                            icon={testResult.passed ? "done" : "clear"}
                            lastItem={true}
                        />
                    </MDBox>
                )
            }))
    }

}
