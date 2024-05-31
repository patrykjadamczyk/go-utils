package base

func Memo1[A0 comparable, R any](f func(A0) R) func(A0) R {
	memoMap := make(map[A0]R)
	return func(a0 A0) R {
		if r, ok := memoMap[a0]; ok {
			return r
		}
		memoMap[a0] = f(a0)
		return memoMap[a0]
	}
}

func Memo2[A0, A1 comparable, R any](f func(A0, A1) R) func(A0, A1) R {
	memoMap := make(map[A0]map[A1]R)
	return func(a0 A0, a1 A1) R {
		if m, ok := memoMap[a0]; ok {
			if r, ok := m[a1]; ok {
				return r
			}
		} else {
			memoMap[a0] = make(map[A1]R)
		}
		memoMap[a0][a1] = f(a0, a1)
		return memoMap[a0][a1]
	}
}

func Memo3[A0, A1, A2 comparable, R any](f func(A0, A1, A2) R) func(A0, A1, A2) R {
	memoMap := make(map[A0]map[A1]map[A2]R)
	return func(a0 A0, a1 A1, a2 A2) R {
		if m, ok := memoMap[a0]; ok {
			if m1, ok := m[a1]; ok {
				if r, ok := m1[a2]; ok {
					return r
				}
			} else {
				memoMap[a0][a1] = make(map[A2]R)
			}
		} else {
			memoMap[a0] = make(map[A1]map[A2]R)
			memoMap[a0][a1] = make(map[A2]R)
		}
		memoMap[a0][a1][a2] = f(a0, a1, a2)
		return memoMap[a0][a1][a2]
	}
}

func Memo4[A0, A1, A2, A3 comparable, R any](f func(A0, A1, A2, A3) R) func(A0, A1, A2, A3) R {
	memoMap := make(map[A0]map[A1]map[A2]map[A3]R)
	return func(a0 A0, a1 A1, a2 A2, a3 A3) R {
		if m, ok := memoMap[a0]; ok {
			if m1, ok := m[a1]; ok {
				if m2, ok := m1[a2]; ok {
					if r, ok := m2[a3]; ok {
						return r
					}
				} else {
					memoMap[a0][a1][a2] = make(map[A3]R)
				}
			} else {
				memoMap[a0][a1] = make(map[A2]map[A3]R)
				memoMap[a0][a1][a2] = make(map[A3]R)
			}
		} else {
			memoMap[a0] = make(map[A1]map[A2]map[A3]R)
			memoMap[a0][a1] = make(map[A2]map[A3]R)
			memoMap[a0][a1][a2] = make(map[A3]R)
		}
		memoMap[a0][a1][a2][a3] = f(a0, a1, a2, a3)
		return memoMap[a0][a1][a2][a3]
	}
}

func Memo5[A0, A1, A2, A3, A4 comparable, R any](f func(A0, A1, A2, A3, A4) R) func(A0, A1, A2, A3, A4) R {
	memoMap := make(map[A0]map[A1]map[A2]map[A3]map[A4]R)
	return func(a0 A0, a1 A1, a2 A2, a3 A3, a4 A4) R {
		if m, ok := memoMap[a0]; ok {
			if m1, ok := m[a1]; ok {
				if m2, ok := m1[a2]; ok {
					if m3, ok := m2[a3]; ok {
						if r, ok := m3[a4]; ok {
							return r
						}
					} else {
						memoMap[a0][a1][a2][a3] = make(map[A4]R)
					}
				} else {
					memoMap[a0][a1][a2] = make(map[A3]map[A4]R)
					memoMap[a0][a1][a2][a3] = make(map[A4]R)
				}
			} else {
				memoMap[a0][a1] = make(map[A2]map[A3]map[A4]R)
				memoMap[a0][a1][a2] = make(map[A3]map[A4]R)
				memoMap[a0][a1][a2][a3] = make(map[A4]R)
			}
		} else {
			memoMap[a0] = make(map[A1]map[A2]map[A3]map[A4]R)
			memoMap[a0][a1] = make(map[A2]map[A3]map[A4]R)
			memoMap[a0][a1][a2] = make(map[A3]map[A4]R)
			memoMap[a0][a1][a2][a3] = make(map[A4]R)
		}
		memoMap[a0][a1][a2][a3][a4] = f(a0, a1, a2, a3, a4)
		return memoMap[a0][a1][a2][a3][a4]
	}
}
