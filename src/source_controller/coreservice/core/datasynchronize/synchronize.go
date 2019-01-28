/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.,
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the ",License",); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an ",AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package instances

import (
	"configcenter/src/common"
	"configcenter/src/common/errors"
	"configcenter/src/common/metadata"
	"configcenter/src/source_controller/coreservice/core"
	"configcenter/src/storage/dal"
)

type synchronizeManager struct {
	dbProxy dal.RDB
}

func (s *synchronizeManager) SynchronizeAdapter(ctx core.ContextParams, syncData *metadata.SynchronizeDataParameter) ([]string, errors.CCError) {
	syncDataAdpater := newSynchronizeDataAdapter(syncData, s.dbProxy)
	err := syncDataAdpater.PreSynchronizeFilter(ctx)
	if err != nil {
		return nil, err
	}
	syncDataAdpater.SaveSynchronize(ctx)
	return syncDataAdpater.GetErrorStringArr(ctx)

}