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
import MDTypography from "components/MDTypography";
import MDAvatar from "components/MDAvatar";

import {Link, useNavigate} from "react-router-dom";


// Images
import TimelineItem from "examples/Timeline/TimelineItem";
import logoGithub from "assets/images/small-logos/github.svg";

export default function buildAssignmentsTable(subject) {

    const Assignment = ({image, name}) => (
        <MDBox display="flex" alignItems="center" lineHeight={1}>
            <MDAvatar src={image} name={name} size="sm"/>
            <MDTypography variant="button" fontWeight="medium" ml={1} lineHeight={1}>
                {name}
            </MDTypography>
        </MDBox>
    );


    const navigate = useNavigate();

    let handleNavigate = (assignment) => {
        navigate("/assignment", {state: assignment})
    }
    if (subject.assignments == null) {
        return {
            columns: [
                {Header: "Assignment", accessor: "companies", width: "45%", align: "left"},
                {Header: "Completion", accessor: "completion", align: "center"},
                {Header: "action", accessor: "action", align: "center"}
            ],
            rows: []
        }
    } else {

        return {
            columns: [
                {Header: "Assignment", accessor: "companies", width: "45%", align: "left"},
                {Header: "Completion", accessor: "completion", align: "center"},
                {Header: "action", accessor: "action", align: "center"}
            ],
            rows: subject.assignments.map((assignment) => (
                {
                    companies: <Assignment image={logoGithub} name={assignment.title}/>,
                    completion: (
                        <MDBox width="2rem" textAlign="right">
                            <TimelineItem
                                color="success"
                                icon="done"
                                lastItem={true}
                            />
                        </MDBox>
                    ),
                    action: (
                        <MDTypography component="a" onClick={() => handleNavigate(assignment)} variant="caption" color="text" fontWeight="medium">
                            Attempt
                        </MDTypography>
                    )
                }))
        }
    }
}
