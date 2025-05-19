// app/components/RegisterPage.js - Register page component
"use client";

import React, { useEffect } from "react";
import { Formik, Form, Field, ErrorMessage } from "formik";
import * as Yup from "yup";
import { useRouter } from "next/navigation";
import { useUser } from "../context/UserContext";

export default function Register() {
    const router = useRouter();
    // get the register function and user data from the user context
    const { register, loggedInUser } = useUser();

    // initial values for the registration form
    const initialValues = {
        username: "",
        email: "",
        password: "",
        confirmPassword: ""
    };

    // validation schema for the registration form
    const validationSchema = Yup.object({
        username: Yup.string().required("Username is required"),
        email: Yup.string()
            .email("Invalid email format")
            .required("Email is required"),
        password: Yup.string()
            .min(6, "Password must be at least 6 characters")
            .required("Password is required"),
        confirmPassword: Yup.string()
            /* validate that the confirmPassword field matches the 
                password field */
            .oneOf([Yup.ref("password"), null], "Passwords must match")
            .required("Confirm Password is required")
    });

    // handle the form submission
    const onSubmit = (values, { resetForm }) => {
        const success = register(values.username, values.email, values.password);

        if (!success) {
            return;
        }

        alert("Account created successfully");
        resetForm();
        router.push("/login");
    };

    // redirect the user to the home page if they are already logged in
    useEffect(() => {
        if (loggedInUser) {
            router.push("/");
            return;
        }
    }, [loggedInUser, router]);
    
    return (
        <div className="register container mt-5">
            <h1 className="text-center mb-4">Register new account</h1>
            <Formik
                initialValues={initialValues}
                validationSchema={validationSchema}
                onSubmit={onSubmit}>
                    {() => (
                        <Form className="card p-4 shadow">
                            <div className="mb-3">
                                <label className="form-label">Username:</label>
                                <Field type="text" 
                                    name="username" 
                                    className="form-control"/>
                                <ErrorMessage 
                                    name="username" 
                                    component="div"
                                    className="text-danger small"/>
                            </div>
                            <div className="mb-3">
                                <label className="form-label">Email:</label>
                                <Field type="email" 
                                    name="email"
                                    className="form-control"/>
                                <ErrorMessage 
                                    name="email" 
                                    component="div"
                                    className="text-danger small"/>
                            </div>
                            <div className="mb-3">
                                <label className="form-label">Password:</label>
                                <Field type="password" 
                                    name="password" 
                                    className="form-control"/>
                                <ErrorMessage 
                                    name="password" 
                                    component="div"
                                    className="text-danger small"/>
                            </div>
                            <div className="mb-3">
                                <label className="form-label">Confirm Password:</label>
                                <Field type="password" 
                                    name="confirmPassword" 
                                    className="form-control"/>
                                <ErrorMessage 
                                    name="confirmPassword" 
                                    component="div"
                                    className="text-danger small"/>
                            </div>
                            <button type="submit" className="btn btn-primary w-100">Register</button>
                        </Form>
                    )}
            </Formik>
        </div>
    );
}
