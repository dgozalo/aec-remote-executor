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


function AssignmentInstructions({assignment}) {
  return (
    <Card id="assignments-card">
      <MDBox pt={1} px={2}>
          <MDTypography variant="h8" fontWeight="medium">
              Assignment: {assignment.title}
          </MDTypography>
        <MDTypography variant="h6" fontWeight="medium">
          Assignment Instructions
        </MDTypography>
      </MDBox>
      <MDBox pt={1} pb={"55%"} px={2}>
        <MDBox component="ul" display="flex" flexDirection="column" p={0} m={0}>
            {assignment.description.split(/\r?\n/).map((line) => (
                <>
                    <MDTypography variant="caption" component="p" color="white" fontWeight={"regular"} >
                        {line}
                    </MDTypography>
                </>
            ))}
            <br/>
            {assignment?.assignment_examples?.map((example) => (
                <>
                    <MDTypography variant="h6" component="b" color="white" fontWeight={"regular"} >
                        {example.title}
                    </MDTypography>
                    {example.description.split(/\r?\n/).map((line) => (
                        <>
                            <MDTypography variant="caption" component="code" color="white" fontWeight={"regular"} >
                                {line}
                            </MDTypography>
                        </>
                    ))}
                    <br/>
                </>
            ))}
        </MDBox>
      </MDBox>
    </Card>

  );
}

export default AssignmentInstructions;
