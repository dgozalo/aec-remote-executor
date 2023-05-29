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
import Grid from "@mui/material/Grid";

// Material Dashboard 2 React components
import MDBox from "components/MDBox";

// Material Dashboard 2 React example components
import DashboardLayout from "examples/LayoutContainers/DashboardLayout";
import DashboardNavbar from "examples/Navbars/DashboardNavbar";
import Footer from "examples/Footer";

// Dashboard components
import AssignmentsDashboardTable from "layouts/dashboard/components/Assignments";
import SubjectsOverview from "layouts/dashboard/components/SubjectsOverview";
import {useState} from "react";

function Dashboard() {
    const [subject, setSubject] = useState("");

    return (
    <DashboardLayout>
      <DashboardNavbar />
      <MDBox py={3}>
        <MDBox>
          <Grid container spacing={3}>
              <Grid item xs={12} md={6} lg={2}>
                  <SubjectsOverview stateChanger={setSubject} />
              </Grid>
            <Grid item xs={12} md={6} lg={10}>
              <AssignmentsDashboardTable subject={subject} />
            </Grid>
          </Grid>
        </MDBox>
      </MDBox>
      <Footer />
    </DashboardLayout>
  );
}

export default Dashboard;
