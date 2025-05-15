// app/components/Login.js - Login component
"use client";

import React, { useEffect } from "react";
import { Formik, Form, Field, ErrorMessage } from "formik";
import * as Yup from "yup";
import { useRouter } from "next/navigation";
import { useUser } from "../context/UserContext";

export default function Login() {
    const router = useRouter();
    const { login, loggedInUser } = useUser();

    const initialValues = {
        email: "",
        password: ""
    };

    const validationSchema = Yup.object({
        email: Yup.string()
            .email("Invalid email format")
            .required("Email is required"),
        password: Yup.string()
            .required("Password is required")
    });
    
    const onSubmit = (values) => {
        const success = login(values.email, values.password);

        if (!success) {
            alert("Invalid email or password");
            return;
        }

        alert("Logged in successfully");
        router.push("/");
    }

    // redirect the user to the home page if already logged in
    useEffect(() => {
        const isLoggedIn = JSON.parse(localStorage.getItem("user"));
        if (isLoggedIn) {
            router.push("/");
            return;
        }
    });

    return (
        <div className="login container mt-5">
            <h1 className="text-center mb-4">Login:</h1>
            <Formik
                initialValues={initialValues}
                validationSchema={validationSchema}
                onSubmit={onSubmit}>
                {() => (
                    <Form className="card p-4 shadow">
                        <div className="mb-3">
                            <label>Email</label>
                            <Field type="email" name="email" className="form-control" />
                            <ErrorMessage 
                                name="email" 
                                component="div" className="text-danger small" />
                        </div>
                        <div className="mb-3">
                            <label>Password</label>
                            <Field type="password" name="password" className="form-control" />
                            <ErrorMessage 
                                name="password" 
                                component="div" className="text-danger small" />
                        </div>
                        <button type="submit" className="btn btn-primary">Log in</button>
                    </Form>
                )}
            </Formik>
        </div>
    );
}
