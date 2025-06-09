'use client';

import React from 'react';
import { useKeycloak } from '../context/KeycloakContext';
import { useUser } from '../context/UserContext';
import { Formik, Form, Field, ErrorMessage } from 'formik';
import * as Yup from 'yup';

export default function CreateServerForm() {
    const { keycloak, authenticated } = useKeycloak();
    const { createServer, isCreating, createError } = useUser();

    const servNameMin= 3;
    const servNameMax= 30;

    const createServerSchema = Yup.object().shape({
        name: Yup.string()
            .trim()
            .required('Server name is required')
            .min(servNameMin, 
                `Server name has to be at least ${servNameMin} characters long`)
            .max(servNameMax, 
                `Server name can't be longer than ${servNameMax} characters`),
    });

    const onSubmit = async (values, { setSubmitting, resetForm, setStatus }) => {
        setStatus(null);

        const server = await createServer(values.name.trim());
        if (server) {
            resetForm();
        } else {
            setStatus(createError || "Failed to create server")
        }
        setSubmitting(false);
    };

    if (!authenticated) {
        return;
    }

    return (
        <div className="mb-4">
            <Formik
                initialValues={{ name: '' }}
                validationSchema={createServerSchema}
                onSubmit={onSubmit}
            >
                {({ isSubmitting, status }) => (
                    <Form>
                        <div className="mb-3">
                            <label>Server Name</label>
                            <Field
                                name="name"
                                type="text"
                                className="form-control"
                                disabled={isSubmitting}
                            />
                            <ErrorMessage 
                                name="name" 
                                component="div" 
                                className="text-danger small"
                            />
                        </div>

                        {status && (
                            <div className="alert alert-danger" role="alert">
                                {status}
                            </div>
                        )}

                        <button
                            type="submit"
                            className="btn btn-primary"
                            disabled={isSubmitting}
                        >
                        {isSubmitting ? (
                            <>
                                <span
                                    className="spinner-border spinner-border-sm me-1"
                                    role="status"
                                    aria-hidden="true"
                                ></span>
                                Creating the server...
                            </>
                        ) : (
                            'Create Server'
                        )}
                        </button>
                    </Form>
                )}
            </Formik>
        </div>
    )
}
