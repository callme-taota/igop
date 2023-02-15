/*
 * Copyright (c) 2022 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package igop

import (
	"errors"
	"fmt"
)

var (
	ErrNoPackage        = errors.New("no package")
	ErrPackage          = errors.New("package contain errors")
	ErrNotFoundMain     = errors.New("not found main package")
	ErrTestFailed       = errors.New("test failed")
	ErrNotFoundPackage  = errors.New("not found package")
	ErrNotFoundImporter = errors.New("not found provider for types.Importer")
	ErrGoexitDeadlock   = errors.New("fatal error: no goroutines (main called runtime.Goexit) - deadlock!")
	ErrNoFunction       = errors.New("no function")
	ErrNoCustomBuiltin  = errors.New("no custom builtin")
	ErrNoTestFiles      = errors.New("[no test files]")
	ErrGoList           = errors.New("error go list")
)

type ExitError int

func (r ExitError) Error() string {
	return fmt.Sprintf("exit %v", int(r))
}
