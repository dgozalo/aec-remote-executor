import React from "react";
import MDTypography from "../../../../components/MDTypography";

const OutputWindow = ({ outputDetails }) => {
    const getOutput = () => {
        let status = outputDetails?.status;

        if (status === 'ERROR') {
            // compilation error
            return (
                <MDTypography variant="caption" component="code" color="red" fontWeight={"regular"} >
                    {outputDetails?.stderr}
                </MDTypography>
            );
        } else if (status === 'COMPLETED') {
            return (
                <MDTypography variant="caption" component="code" color="red" fontWeight={"regular"} >
                    {outputDetails.stdout !== null  ? `${outputDetails.stdout}` : null}
                </MDTypography>
            );
        }
    };
    return (
        <>
            <MDTypography variant="h3" color="white" component="p" color="white"  >
                Output
            </MDTypography>

            <MDTypography variant="caption" component="p" color="white"  fontWeight={"regular"}>
                {outputDetails ? <>{getOutput()}</> : null}
            </MDTypography>
        </>
    );
};

export default OutputWindow;