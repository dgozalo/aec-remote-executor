import React from "react";
import MDTypography from "../../../../components/MDTypography";
import Grid from "@mui/material/Grid";
import DataTable from "examples/Tables/DataTable";
import buildTestResultsTable from "./data";

const OutputDetails = ({outputDetails}) => {
    return (
        <Grid item >
            <MDTypography variant={"caption"} component={"p"} color={"white"} fontWeight={"regular"}>
                Status:{" "}
                {outputDetails?.status}
            </MDTypography>
            { outputDetails?.testResults &&
                    <DataTable
                        table={buildTestResultsTable(outputDetails?.testResults)}
                        showTotalEntries={false}
                        isSorted={false}
                        noEndBorder
                        entriesPerPage={false}
                    />
                }
        </Grid>
    );
};

export default OutputDetails;