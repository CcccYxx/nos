/*
 * Copyright 2023 nebuly.com.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gpu

type Model string

func (m Model) String() string {
	return string(m)
}

const (
	GPUModel_A30            Model = "A30"
	GPUModel_A100_SXM4_40GB Model = "NVIDIA-A100-40GB-SXM4"
	GPUModel_A100_PCIe_80GB Model = "NVIDIA-A100-80GB-PCIe"
	GPUModel_H100_SXM5_96GB Model = "NVIDIA-H100-96GB-SXM5"
	GPUModel_H100_PCIe_96GB Model = "NVIDIA-H100-96GB-PCIe"
)
