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

// Material Dashboard 2 React examples
import DashboardLayout from "examples/LayoutContainers/DashboardLayout";
import DashboardNavbar from "examples/Navbars/DashboardNavbar";
import Footer from "examples/Footer";
import MasterCard from "examples/Cards/MasterCard";
import DefaultInfoCard from "examples/Cards/InfoCards/DefaultInfoCard";

// AssignmentInstructions page components
import AssignmentInstructions from "./components/AssignmentInstructions";
import CodeEditor from "../../examples/Custom/CodeEditor";
import {useLocation, useNavigate} from "react-router-dom";

function Assignments() {
    const location = useLocation();
    const navigate = useNavigate();

    //if the state is null, then the user has not selected an assignment and should be redirected to the dashboard
    if (location.state == null) {
        navigate('/dashboard', {replace: true});
    }
    console.log(location.state);

  return (
    <DashboardLayout>
      <DashboardNavbar absolute isMini />
      <MDBox mt={8}>
        <MDBox mb={3}>
          <Grid container spacing={3}>
              <Grid item xs={5} md={5}>
                  <AssignmentInstructions assignment={location.state} />
              </Grid>
              <Grid item xs={5} md={7}>
                  <CodeEditor assignment={location.state}/>
              </Grid>
          </Grid>
        </MDBox>
      </MDBox>
      <Footer />
    </DashboardLayout>
  );
}

export default Assignments;
