// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

func init() {
	// Register defaulting functions until the following upstream issue is resolved:
	//
	// https://github.com/kubernetes/kubernetes/issues/135417#issuecomment-3655543270
	//
	// TODO: Remove this one after the issue above is resolved.
	localSchemeBuilder.Register(RegisterDefaults)
}
