import Editor from "@monaco-editor/react";
import {useEffect, useState} from "react";
import LanguagesDropdown from "./components/LanguageDropdown";
import ThemeDropdown from "./components/ThemeDropdown";
import {toast} from "react-toastify";
import Grid from "@mui/material/Grid";
import MDBox from "../../../components/MDBox";
import OutputWindow from "./components/OutputWindow";
import {languageOptions} from "./constants/languageOptions";
import useKeyPress from "./useKeyPress";
import {defineTheme} from "./lib/defineTheme";
import OutputDetails from "./components/OutputDetails";
import MDTypography from "../../../components/MDTypography";
import {gql, useLazyQuery, useMutation} from "@apollo/client";
import Card from "@mui/material/Card";

const EXECUTION_MUTATION = gql`
    mutation runExecution($input:NewExecution!) {
      runExecution(input:$input) {
        id,
        language,
        code
      }
    }
`;

const EXECUTION_QUERY = gql`
query getExecutionStatus($id:ID!) {
  GetExecutionStatus(id:$id) {
    stdout,
    stderr,
    status,
    testResults {
      testName,
      expected,
      actual,
      passed
    }
  }
}
`;

function CodeEditor({assignment, setOutputDetails, outputDetails}) {
    const selectedLanguage = assignment?.assignment_code_templates[0];
    const [code, setCode] = useState("");
    const [processing, setProcessing] = useState(null);
    const [theme, setTheme] = useState("cobalt");
    const [language, setLanguage] = useState(null);
    const [executeCode] = useMutation(EXECUTION_MUTATION);
    const [runQuery, { called, loading, data }] = useLazyQuery(EXECUTION_QUERY)
    const enterPress = useKeyPress("Enter");
    const ctrlPress = useKeyPress("Control");

    useEffect(() => {
        languageOptions.filter((lang) => {
            if (lang.value === selectedLanguage.language) {
                setLanguage(lang);
                setCode(selectedLanguage.code);
            }
        });
    }, [selectedLanguage]);

    const onSelectChange = (sl) => {
        assignment.assignment_code_templates.filter((template) => {
            if (template.language === sl.value) {
                setCode(template.code);
                setLanguage(sl);
            }
        });
    };

    useEffect(() => {
        if (enterPress && ctrlPress) {
            handleCompile();
        }
    }, [ctrlPress, enterPress]);

    const onChange = (action) => {
        console.log("onChange", action, JSON.stringify(action, null, 2));
        setCode(action);
    }
    const handleCompile = () => {
        setProcessing(true)

        executeCode({
            variables: {
                input: {
                    language: language.value,
                    code: code,
                    assignmentId: assignment.id
                }
            }
        }).then((res) => {
            checkStatus(res.data.runExecution.id)
        }).catch((err) => {
            console.log("Error");
            console.log(err)
        });
    };

    const checkStatus = (executionId) => {
        // We will come to the implementation later in the code
        runQuery({variables: {id: executionId}, fetchPolicy: "network-only"}).then((res) => {
            if (res.data.GetExecutionStatus.status === "RUNNING") {
                setTimeout(() => {
                    checkStatus(executionId)
                }, 1000)
            } else {
                setOutputDetails({
                    status: res.data.GetExecutionStatus.status,
                    stdout: res.data.GetExecutionStatus.stdout,
                    stderr: res.data.GetExecutionStatus.stderr,
                    testResults: res.data.GetExecutionStatus.testResults,
                    memory: "1gb",
                    time: "1s",
                });
                setProcessing(false)
            }
        });

    };

    function handleThemeChange(th) {
        const theme = th;
        if (["light", "vs-dark"].includes(theme.value)) {
            setTheme(theme);
        } else {
            defineTheme(theme.value).then((_) => setTheme(theme));
        }
    }

    useEffect(() => {
        defineTheme("oceanic-next").then((_) =>
            setTheme({value: "oceanic-next", label: "Oceanic Next"})
        );
    }, []);

    const showSuccessToast = (msg) => {
        toast.success(msg || `Compiled Successfully!`, {
            position: "top-right",
            autoClose: 1000,
            hideProgressBar: false,
            closeOnClick: true,
            pauseOnHover: true,
            draggable: true,
            progress: undefined,
        });
    };
    const showErrorToast = (msg) => {
        toast.error(msg || `Something went wrong! Please try again.`, {
            position: "top-right",
            autoClose: 1000,
            hideProgressBar: false,
            closeOnClick: true,
            pauseOnHover: true,
            draggable: true,
            progress: undefined,
        });
    };
    const classnames = (...args) => {
        return args.join(" ");
    };
    const alertContent = (name) => (
        <MDTypography variant="body2" color="white">
            A simple {name} alert with{" "}
            <MDTypography component="a" href="#" variant="body2" fontWeight="medium" color="white">
                an example link
            </MDTypography>
            . Give it a click if you like.
        </MDTypography>
    );
    return (
        <grid container spacing={3}>
            <Grid >
                <MDBox mb={1}>
                    <Grid container spacing={1}>
                        <Grid item xs={5} md={3}>
                            <LanguagesDropdown onSelectChange={onSelectChange}/>
                        </Grid>
                        <Grid item xs={5} md={7}>
                            <ThemeDropdown handleThemeChange={handleThemeChange} theme={theme}/>
                        </Grid>
                        <Grid item xs={5} md={25}>
                            <Editor
                                height={`45vh`}
                                width={`100%`}
                                language={language?.value}
                                theme={theme.value}
                                defaultValue={code}
                                value={code}
                                onChange={onChange}
                            />
                        </Grid>
                        <Grid m={1}>
                            <Grid item mb={1}>
                                <button
                                    onClick={handleCompile}
                                    disabled={!code}
                                    className={classnames(
                                        "mt-4 border-2 border-black z-10 rounded-md shadow-[5px_5px_0px_0px_rgba(0,0,0)] px-4 py-2 hover:shadow transition duration-200 bg-white flex-shrink-0",
                                        !code ? "opacity-50" : ""
                                    )}
                                >
                                    {processing ? "Processing..." : "Compile and Execute"}
                                </button>
                            </Grid>
                            <Grid item width={"97vh"} height={"30vh "}>
                                <Card container id="assignments-card">
                                    <MDBox pt={1} px={3} >
                                        <Grid item mb={4.5}>
                                            <OutputWindow outputDetails={outputDetails}/>
                                            {outputDetails && <OutputDetails outputDetails={outputDetails}/>}
                                        </Grid>
                                    </MDBox>
                                </Card>
                            </Grid>
                        </Grid>
                    </Grid>
                </MDBox>
            </Grid>
        </grid>

    )
}

export default CodeEditor;