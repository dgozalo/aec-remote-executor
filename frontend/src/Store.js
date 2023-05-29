import {configureStore, createSlice} from "@reduxjs/toolkit";

export const subjectsSlice = createSlice({
    name: "subjects",
    initialState: {
        value: "",
    },
    reducers: {
        changeSubject: (state, action) => {
            state.value = action.payload;
        }
    }
});

export const { changeSubject } = subjectsSlice.actions

export const store = configureStore({
    reducer: {
        subjects: subjectsSlice.reducer,
    }
})