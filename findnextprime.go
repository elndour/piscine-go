package piscine

func FindNextPrime(nb int) int {
	if nb <= 0 {
		return 2
	} else {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				for i := 2; i < nb; i++ {
					if (nb+1)%i == 0 {
						for i := 2; i < nb; i++ {
							if (nb+2)%i == 0 {
								for i := 2; i < nb; i++ {
									if (nb+3)%i == 0 {
										for i := 2; i < nb; i++ {
											if (nb+4)%i == 0 {
												for i := 2; i < nb; i++ {
													if (nb+5)%i == 0 {
														for i := 2; i < nb; i++ {
															if (nb+6)%i == 0 {
																for i := 2; i < nb; i++ {
																	if (nb+7)%i == 0 {
																		return nb + 8
																	}
																}
																return nb + 7
															}
														}
														return nb + 6
													}
												}
												return nb + 5
											}
										}
										return nb + 4
									}
								}
								return nb + 3
							}
						}
						return nb + 2
					}
				}
				return nb + 1
			}
		}
		return nb
	}
}
