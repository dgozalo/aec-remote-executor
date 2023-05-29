import React from "react";
import MDTypography from "../../../../components/MDTypography";

const OutputDetails = ({outputDetails}) => {
    return (
        <div className="metrics-container mt-4 flex flex-col space-y-3">
            <MDTypography variant={"caption"} component={"p"} color={"white"} fontWeight={"regular"}>
                Status:{" "}
                {outputDetails?.status?.description}
            </MDTypography>
            <MDTypography variant={"caption"} component={"p"} color={"white"} fontWeight={"regular"}>
                Memory:{" "}
                {outputDetails?.memory}
            </MDTypography>
            <MDTypography variant={"caption"} component={"p"} color={"white"} fontWeight={"regular"}>
                Time:{" "}
                {outputDetails?.time}
            </MDTypography>
        </div>
    );
};

export default OutputDetails;