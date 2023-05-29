import React from "react";
import MDTypography from "../../../../components/MDTypography";

const OutputWindow = ({ outputDetails }) => {
    const getOutput = () => {
        let statusId = outputDetails?.status?.id;

        if (statusId === 6) {
            // compilation error
            return (
                <pre className="px-2 py-1 font-normal text-xs text-red-500">
          {outputDetails?.compile_output}
        </pre>
            );
        } else if (statusId === 3) {
            return (
                <pre className="px-2 py-1 font-normal text-xs text-green-500">
          {outputDetails.stdout !== null
              ? `${outputDetails.stdout}`
              : null}
        </pre>
            );
        } else if (statusId === 5) {
            return (
                <pre className="px-2 py-1 font-normal text-xs text-red-500">
          {`Time Limit Exceeded`}
        </pre>
            );
        } else {
            return (
                <pre className="px-2 py-1 font-normal text-xs text-red-500">
          {outputDetails?.stderr}
        </pre>
            );
        }
    };
    return (
        <>
            <MDTypography variant="h3" color="white" component="p" color="white"  >
                Output
            </MDTypography>
            <MDTypography variant="caption" component="code" color="white"  fontWeight={"regular"}>
                {outputDetails ? <>{getOutput()}</> : null}
            </MDTypography>
        </>
    );
};

export default OutputWindow;