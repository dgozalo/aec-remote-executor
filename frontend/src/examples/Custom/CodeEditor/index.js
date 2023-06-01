import Editor from "@monaco-editor/react";
import {useEffect, useState} from "react";
import LanguagesDropdown from "./components/LanguageDropdown";
import ThemeDropdown from "./components/ThemeDropdown";
import {toast, ToastContainer} from "react-toastify";
import Grid from "@mui/material/Grid";
import MDBox from "../../../components/MDBox";
import OutputWindow from "./components/OutputWindow";
import {languageOptions} from "./constants/languageOptions";
import useKeyPress from "./useKeyPress";
import {defineTheme} from "./lib/defineTheme";
import OutputDetails from "./components/OutputDetails";
import MDAlert from "../../../components/MDAlert";
import MDTypography from "../../../components/MDTypography";

const javascriptDefault = `// some comment`;

function CodeEditor() {
    const [code, setCode] = useState(javascriptDefault);
    const [customInput, setCustomInput] = useState("");
    const [outputDetails, setOutputDetails] = useState(null);
    const [processing, setProcessing] = useState(null);
    const [theme, setTheme] = useState("cobalt");
    const [language, setLanguage] = useState(languageOptions[0]);

    const enterPress = useKeyPress("Enter");
    const ctrlPress = useKeyPress("Control");

    const onSelectChange = (sl) => {
        setLanguage(sl);
    };

    useEffect(() => {
        if (enterPress && ctrlPress) {
            handleCompile();
        }
    }, [ctrlPress, enterPress]);

    const onChange = (action) => {
        setCode(action);
    }
    const handleCompile = () => {
        setProcessing(true)
        setOutputDetails({
            status: {
                id: 5,
                message: "Compiling...",
            },
            compile_output: "asasasas",
            stdout: "eeeee",
            stderr: "cercrrrwe",
            memory: "1gb",
            time: "1s",
        })

        console.log("Compiling " + language.label + " code..." + code);
        setTimeout(() => {
            setProcessing(false)
            showSuccessToast("Compiled Successfully!")
        }, 2000)
    };

    const checkStatus = async (token) => {
        // We will come to the implementation later in the code
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
            setTheme({ value: "oceanic-next", label: "Oceanic Next" })
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
        <>
            {outputDetails && <MDAlert color="secondary" dismissible>
                {alertContent("secondary")}
            </MDAlert>}
            <MDBox>
                <MDBox mb={1}>
                    <Grid container spacing={0}>
                        <Grid item xs={5} md={5}>
                            <LanguagesDropdown onSelectChange={onSelectChange} />
                        </Grid>
                        <Grid item xs={5} md={7}>
                            <ThemeDropdown handleThemeChange={handleThemeChange} theme={theme} />
                        </Grid>
                        <Grid item xs={5} md={25}>
                            <Editor
                                height={`62vh`}
                                width={`100%`}
                                language={language?.value}
                                code={code}
                                theme={theme.value}
                                defaultValue="// some comment"
                                onChange={onChange}
                            />
                        </Grid>
                        <Grid >
                            <OutputWindow outputDetails={outputDetails} />
                            <Grid >
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
                            {outputDetails && <OutputDetails outputDetails={outputDetails} />}
                        </Grid>
                    </Grid>
                </MDBox>
            </MDBox>
        </>

            )
}

export default CodeEditor;