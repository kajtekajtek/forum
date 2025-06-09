"use client";
import React from "react";
import { useKeycloak } from "../context/KeycloakContext";
import { Formik, Form, Field, ErrorMessage } from "formik";
import * as Yup from "yup";
import { createChannel as apiCreateChannel } from "../../lib/api/apiClient";

export default function CreateChannelForm({ serverId, onCreated }) {
    const { keycloak, authenticated } = useKeycloak();

    const nameMin = 1;
    const nameMax = 30;

    const schema = Yup.object().shape({
        name: Yup.string()
            .trim()
            .required("Channel name is required")
            .min(nameMin, `Channel name has to be at least ${nameMin} characters long`)
            .max(nameMax, `Channel name can't be longer than ${nameMax} characters`)
    });

    const onSubmit = async (values, { setSubmitting, resetForm, setStatus }) => {
        setStatus(null);
        try {
            const channel = await apiCreateChannel(
                keycloak.token,
                serverId,
                values.name.trim()
            );
            resetForm();
            if (onCreated) onCreated(channel);
        } catch (err) {
            console.error(err);
            setStatus("Failed to create channel");
        } finally {
            setSubmitting(false);
        }
    };

    if (!authenticated) return null;

    return (
        <div className="mb-3">
            <Formik initialValues={{ name: "" }} validationSchema={schema} onSubmit={onSubmit}>
                {({ isSubmitting, status }) => (
                    <Form>
                        <div className="mb-3">
                            <label>Channel Name</label>
                            <Field
                                name="name"
                                type="text"
                                className="form-control"
                                placeholder="Channel Name"
                                disabled={isSubmitting}
                            />
                            <ErrorMessage name="name" component="div" className="text-danger small" />
                        </div>
                        {status && (
                            <div className="alert alert-danger" role="alert">
                                {status}
                            </div>
                        )}
                        <button type="submit" className="btn btn-primary" disabled={isSubmitting}>
                            {isSubmitting ? (
                            <>
                                <span className="spinner-border spinner-border-sm me-1" role="status" aria-hidden="true"></span>
                                Creating channel...
                            </>
                            ) : (
                                "Create Channel"
                            )}
                        </button>
                    </Form>
                )}
            </Formik>
        </div>
    );
}
