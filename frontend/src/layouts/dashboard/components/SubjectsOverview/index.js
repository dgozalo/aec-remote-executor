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

// Material Dashboard 2 React example components
import TimelineItem from "examples/Timeline/TimelineItem";
import {gql, useQuery} from "@apollo/client";

//TODO: Parametrize this query
const ALUM_QUERY = gql`
  {
    GetAlumnus(id: "1"){
        id,
        first_name,
        last_name,
        email,
        graduation_year,
        subjects {
          id,
          name,
          semester,
          assignments {
            id,
            title,
            description,
            assignment_examples {
              id,
              title,
              description
            },
            assignment_code_templates {
              id,
              language,
              code
            }
          }
        }
    }
  }
`;

function SubjectsOverview({stateChanger, ...rest}) {
    const { data, loading, error } = useQuery(ALUM_QUERY);
    if (loading) return "Loading...";
    if (error) return <pre>{error.message}</pre>
    return (
        <Card sx={{height: "100%"}}>
            <MDBox pt={3} px={3}>
                <MDTypography variant="h6" fontWeight="medium">
                    Subjects overview
                </MDTypography>
            </MDBox>
            <MDBox p={2}>
                {data.GetAlumnus.subjects.map((subject) => (
                    <MDBox p={0} onClick={() => {
                        stateChanger(subject);
                    }}>
                        <TimelineItem
                            color="light"
                            icon="auto_stories"
                            title={subject.name}
                            id={subject.id}
                            lastItem={true}
                        />
                    </MDBox>
                ))}
            </MDBox>
        </Card>
    );
}

export default SubjectsOverview;
