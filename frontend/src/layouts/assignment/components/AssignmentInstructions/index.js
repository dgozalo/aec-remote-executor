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
import Card from "@mui/material/Card";

// Material Dashboard 2 React components
import MDBox from "components/MDBox";
import MDTypography from "components/MDTypography";


function AssignmentInstructions() {

  return (
    <Card id="delete-account">
      <MDBox pt={3} px={2}>
          <MDTypography variant="h8" fontWeight="medium">
              Assignment: Two Sum
          </MDTypography>
        <MDTypography variant="h6" fontWeight="medium">
          Assignment Instructions
        </MDTypography>
      </MDBox>
      <MDBox pt={1} pb={2} px={2}>
        <MDBox component="ul" display="flex" flexDirection="column" p={0} m={0}>
            <MDTypography variant="caption" component="p" color="white" fontWeight={"regular"} >
                    Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
            </MDTypography>
                <br/>
            <MDTypography variant="caption" component="p" color="white" fontWeight={"regular"} >
                You may assume that each input would have exactly one solution, and you may not use the same element twice.
            </MDTypography>
                <br/>
            <MDTypography variant="caption" component="p" color="white" fontWeight={"regular"} >
                    You can return the answer in any order.
            </MDTypography>
                <br/>
            <MDTypography variant="h6" color="white" component="b" color="white"  >
                    Example 1:
            </MDTypography>
            <MDTypography variant="caption" color="white" component="code" color="white"  >
                    Input: nums = [2,7,11,15], target = 9
            </MDTypography>
            <MDTypography variant="caption" color="white" component="code" color="white"  >
                    Output: [0,1]
            </MDTypography>
            <MDTypography variant="caption" color="white" component="code" color="white"  >
                    Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
            </MDTypography>
            <br/>
            <MDTypography variant="h6" color="white" component="b" color="white"  >
                Example 1:
            </MDTypography>
            <MDTypography variant="caption" color="white" component="code" color="white"  >
                Input: nums = [2,7,11,15], target = 9
            </MDTypography>
            <MDTypography variant="caption" color="white" component="code" color="white"  >
                Output: [0,1]
            </MDTypography>
            <MDTypography variant="caption" color="white" component="code" color="white"  >
                Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
            </MDTypography>
            <br/>
            <MDTypography variant="h6" color="white" component="b" color="white"  >
                Example 1:
            </MDTypography>
            <MDTypography variant="caption" color="white" component="code" color="white"  >
                Input: nums = [2,7,11,15], target = 9
            </MDTypography>
            <MDTypography variant="caption" color="white" component="code" color="white"  >
                Output: [0,1]
            </MDTypography>
            <MDTypography variant="caption" color="white" component="code" color="white"  >
                Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
            </MDTypography>
            <br/>
            <MDTypography variant="caption" component="p" color="white" fontWeight={"regular"}  >
Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
            </MDTypography>
        </MDBox>
      </MDBox>
    </Card>

  );
}

export default AssignmentInstructions;
